package internal

import (
	"bufio"
	"context"
	"errors"
	"io"
	"net/http"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/google/go-github/v57/github"
)

type GitDiff struct {
	// FilePathOld represents the old file path in the diff, typically
	// indicated by a line starting with "---". This is the file path
	// before the changes were made.
	FilePathOld string

	// FilePathNew represents the new file path in the diff, typically
	// indicated by a line starting with "+++ ". This is the file path
	// after the changes were made. In most cases, it is the same as
	// FilePathOld unless the file was renamed or moved.
	FilePathNew string

	// Index is a string that usually contains the hash values before
	// and after the changes, along with some additional metadata.
	// This line typically starts with "index" in the diff output.
	Index string

	// DiffContents contains the actual content of the diff. This part
	// of the struct includes the changes made to the file, typically
	// represented by lines starting with "+" (additions) or "-"
	// (deletions). It includes all the lines that show the modifications
	// to the file.
	DiffContents string
}

// ParseGithubPullRequestURL parses a GitHub pull request URL and returns the owner, repository,
// and pull request number.
func ParseGithubPullRequestURL(pullRequestURL string) (string, string, string, error) {
	parts := strings.Split(pullRequestURL, "/")

	if len(parts) != 7 {
		return "", "", "", errors.New("invalid pull request URL")
	}

	owner := parts[3]
	repo := parts[4]
	prNumber := parts[6]

	return owner, repo, prNumber, nil
}

func GetPullRequest(owner string, repo string, gptModel string, prNumber int, callback OnGetPullRequestEvent) error {
	timeout, err := GetContextTimeout()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	client := github.NewClient(nil)
	pullRequest, _, err := client.PullRequests.Get(ctx, owner, repo, prNumber)
	if err != nil {
		return err
	}

	pullRequest.GetDiffURL()

	ch := make(chan string)
	go getDiffContents(ch, pullRequest.GetDiffURL())

	contents := <-ch

	callback(contents, gptModel)

	return nil
}

func getDiffContents(c chan<- string, diffURL string) {
	diffContents, err := http.Get(diffURL)
	if err != nil {
		c <- err.Error()
		return
	}

	bodyBytes, err := io.ReadAll(diffContents.Body)
	if err != nil {
		c <- err.Error()
		return
	}

	c <- string(bodyBytes)
}

// ParseGitDiff takes a string representing a combined Git diff and a list of
// file extensions to ignore. It returns a slice of GitDiff structs, each representing
// a parsed file diff. The function performs the following steps:
//  1. Splits the combined Git diff into individual file diffs using the
//     splitDiffIntoFiles function. This function looks for "diff --git" as a
//     delimiter to separate each file's diff.
//  2. Iterates over each file diff string. For each string, it:
//     a. Attempts to parse the string into a GitDiff struct using the
//     parseGitDiffFileString function. This function extracts the old and new
//     file paths, index information, and the actual diff content.
//     b. Checks for parsing errors. If an error occurs, it skips the current file
//     diff and continues with the next one.
//  3. Filters out file diffs based on the provided ignore list. The ignore list
//     contains file extensions (e.g., ".mod"). The function uses the
//     getFileExtension helper to extract the file extension from the new file path
//     (FilePathNew) of each GitDiff struct. If the extension matches any in the
//     ignore list, the file diff is skipped.
//  4. Appends the successfully parsed and non-ignored GitDiff structs to the
//     filteredList slice.
//
// Parameters:
//   - diff: A string representing the combined Git diff.
//   - ignoreList: A slice of strings representing the file extensions to ignore.
//
// Returns:
//   - A slice of GitDiff structs, each representing a parsed and non-ignored file diff.
func ParseGitDiff(diff string, ignoreList []string) []*GitDiff {
	files := splitDiffIntoFiles(diff)
	var filteredList []*GitDiff

	for _, file := range files {

		gitDiff, err := parseGitDiffFileString(file)
		if err != nil {
			continue
		}

		if slices.Contains(ignoreList, getFileExtension(gitDiff.FilePathNew)) {
			continue
		}

		filteredList = append(filteredList, gitDiff)
	}

	return filteredList
}

// splitDiffIntoFiles splits a single diff string into a slice of
// strings, where each string represents the diff of an individual file.
// It assumes that 'diff --git' is used as a delimiter between file diffs.
func splitDiffIntoFiles(diff string) []string {
	var files []string
	var curFile strings.Builder

	scanner := bufio.NewScanner(strings.NewReader(diff))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "diff --git") {
			// Detected start of new file
			if curFile.Len() > 0 {
				files = append(files, strings.TrimSpace(curFile.String()))
				curFile.Reset()
			}
			curFile.WriteString(line + "\n")
		} else {
			curFile.WriteString(line + "\n")
		}
	}

	// Add the last file diff to the list
	if curFile.Len() > 0 {
		files = append(files, strings.TrimSpace(curFile.String()))
	}

	return files
}

// ParseGitDiffFileString takes a string input representing a Git diff of a single file
// and returns a GitDiff struct containing the parsed information. The input
// string is expected to contain at least four lines, including the file paths
// line, the index line, and the diff content. The function performs the following
// steps to parse the diff:
//  1. Splits the input string into lines.
//  2. Validates that there are enough lines to form a valid Git diff.
//  3. Extracts the old and new file paths from the first line. The line is
//     expected to contain two file paths separated by a space.
//  4. Extracts the index information from the second line. The line should
//     start with "index " followed by the index information.
//  5. Joins the remaining lines, starting from the third line, to form the
//     diff content.
//
// The function returns an error if the input is not in the expected format,
// such as if there are not enough lines, if the file paths line is invalid,
// or if the index line is incorrectly formatted.
//
// Parameters:
//   - input: A string representing the Git diff of a single file.
//
// Returns:
//   - A pointer to a GitDiff struct containing the parsed file paths, index,
//     and diff content.
//   - An error if the input string is not in the expected format or if any
//     parsing step fails.
func parseGitDiffFileString(input string) (*GitDiff, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	var (
		filePaths []string
		index     string
		diff      []string
	)

	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case strings.HasPrefix(line, "diff --git"):
			filePaths = strings.Fields(line)[2:]
			if len(filePaths) != 2 {
				return nil, errors.New("invalid file paths")
			}
		case strings.HasPrefix(line, "index "):
			index = strings.TrimSpace(line[6:])
		default:
			diff = append(diff, line)
		}
	}

	if len(filePaths) == 0 || len(index) == 0 || len(diff) == 0 {
		return nil, errors.New("invalid git diff format")
	}

	return &GitDiff{
		FilePathOld:  filePaths[0],
		FilePathNew:  filePaths[1],
		Index:        index,
		DiffContents: strings.Join(diff, "\n"),
	}, nil
}

func getFileExtension(path string) string {
	// If the path ends with a slash, it's a directory; return an empty string
	if strings.HasSuffix(path, string(filepath.Separator)) {
		return ""
	}

	fileName := filepath.Base(path)

	// Check if the path is a directory or empty
	if fileName == "." || fileName == "/" || fileName == "" {
		return ""
	}

	// Check for dot files (hidden files in Unix-based systems)
	if len(fileName) > 1 && fileName[0] == '.' && strings.Count(fileName, ".") == 1 {
		return fileName
	}

	// Extract the extension
	return filepath.Ext(fileName)
}

func GetContextTimeout() (time.Duration, error) {
	cfg, err := GetConfig()
	if err != nil {
		return 0, err
	}

	return cfg.ContextTimeout, nil
}
