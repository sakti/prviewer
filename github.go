package main

import (
	"bytes"
	"context"
	"fmt"

	"github.com/google/go-github/v24/github"
	"golang.org/x/oauth2"

	"github.com/andlabs/ui"
)

func getPR(base string) {
	buf := new(bytes.Buffer)

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cfg.GithubToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	opt := &github.PullRequestListOptions{
		State: "all",
		Base:  base,
	}

	prs, _, err := client.PullRequests.List(ctx, cfg.RepoOwner, cfg.RepoName, opt)
	if err != nil {
		Error(err.Error())
	}
	Info("get list of PR succeed")

	ui.QueueMain(func() {
		progressProcess.SetValue(70)
	})

	optComment := &github.IssueListCommentsOptions{
		Sort: "created",
	}

	for _, pr := range prs {
		buf.WriteString(fmt.Sprintf("PR #%d, %s: %s\n", pr.GetNumber(), pr.GetState(), pr.GetTitle()))
		buf.WriteString(pr.GetBody())
		buf.Write([]byte("\n"))

		buf.WriteString("Comments:\n")

		commentCtx := context.Background()
		comments, _, err := client.Issues.ListComments(commentCtx, cfg.RepoOwner, cfg.RepoName, pr.GetNumber(), optComment)
		if err != nil {
			Error(err.Error())
			continue
		}
		if len(comments) == 0 {
			buf.WriteString("no comment\n")
		}
		for _, comment := range comments {
			buf.WriteString(comment.GetUser().GetLogin() + ":\n")
			buf.WriteString(comment.GetBody() + "\n")
		}

		buf.Write([]byte("\n"))
		buf.Write([]byte("\n"))
		buf.Write([]byte("\n"))
		buf.WriteString("================================================================================")
		buf.Write([]byte("\n"))
	}

	ui.QueueMain(func() {
		progressProcess.SetValue(100)
	})

	ui.QueueMain(func() {
		btnGetPR.Enable()
		comboBranch.Enable()
		entryPR.SetText(buf.String())
	})
}
