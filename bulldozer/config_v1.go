// Copyright 2018 Palantir Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bulldozer

type MessageStrategy string
type TitleStrategy string
type MergeMethod string

const (
	PullRequestBody  MessageStrategy = "pull_request_body"
	SummarizeCommits MessageStrategy = "summarize_commits"
	EmptyBody        MessageStrategy = "empty_body"

	PullRequestTitle   TitleStrategy = "pull_request_title"
	FirstCommitTitle   TitleStrategy = "first_commit_title"
	GithubDefaultTitle TitleStrategy = "github_default"

	MergeCommit     MergeMethod = "merge"
	SquashAndMerge  MergeMethod = "squash"
	RebaseAndMerge  MergeMethod = "rebase"
	FastForwardOnly MergeMethod = "ff-only"
)

type MergeConfig struct {
	Trigger Signals `yaml:"trigger"`
	Ignore  Signals `yaml:"ignore"`

	// Blacklist and Whitelist are legacy options that will be disabled in a future v2 format
	Blacklist Signals `yaml:"blacklist"`
	Whitelist Signals `yaml:"whitelist"`

	DeleteAfterMerge       bool `yaml:"delete_after_merge"`
	AllowMergeWithNoChecks bool `yaml:"allow_merge_with_no_checks"`

	Method       MergeMethod              `yaml:"method"`
	MergeMethods []ConditionalMergeMethod `yaml:"merge_method"`
	Options      MergeOptions             `yaml:"options"`

	BranchMethod map[string]MergeMethod `yaml:"branch_method"`

	// Additional status checks that bulldozer should require
	// (even if the branch protection settings doesn't require it)
	RequiredStatuses []string `yaml:"required_statuses"`
}

type MergeOptions struct {
	Squash *SquashOptions `yaml:"squash"`
}

type ConditionalMergeMethod struct {
	Method  MergeMethod `yaml:"method"`
	Trigger Signals     `yaml:"trigger"`
}

type SquashOptions struct {
	Title            TitleStrategy   `yaml:"title"`
	Body             MessageStrategy `yaml:"body"`
	MessageDelimiter string          `yaml:"message_delimiter"`
}

type UpdateConfig struct {
	Trigger Signals `yaml:"trigger"`
	Ignore  Signals `yaml:"ignore"`

	IgnoreDrafts *bool `yaml:"ignore_drafts"`

	// Blacklist and Whitelist are legacy options that will be disabled in a future v2 format
	Blacklist Signals `yaml:"blacklist"`
	Whitelist Signals `yaml:"whitelist"`
}

type Config struct {
	Version int `yaml:"version"`

	Merge  MergeConfig  `yaml:"merge"`
	Update UpdateConfig `yaml:"update"`
}
