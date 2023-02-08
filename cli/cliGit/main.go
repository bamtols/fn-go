package main

import (
	"fmt"
	"github.com/bamtols/fn-go/fn/fnPanic"
	"github.com/spf13/cobra"
	"os/exec"
)

type ()

func main() {
	cmd := &cobra.Command{
		Use: "cliGit",
	}

	cmd.AddCommand(saveUp())

	fnPanic.HasError(cmd.Execute())
}

func saveUp() (root *cobra.Command) {
	root = &cobra.Command{
		Use: "save",
	}

	fileNm := "tags.yaml"
	fileNmFlagNm := "file"
	commitMsg := ""
	commitMsgFlagNm := "msg"

	root.Flags().StringVar(&fileNm, fileNmFlagNm, fileNm, "filename")
	root.Flags().StringVar(&commitMsg, commitMsgFlagNm, commitMsg, "commit message")

	root.Run = func(cmd *cobra.Command, args []string) {
		fileNm = fnPanic.HasErrorOrValue(cmd.Flags().GetString(fileNmFlagNm))
		commitMsg = fnPanic.HasErrorOrValue(cmd.Flags().GetString(commitMsgFlagNm))

		mng := NewFileMng(fileNm)

		data := fnPanic.HasErrorOrValue(mng.Open())

		gitBranch := fnPanic.HasErrorOrValue(mng.GetGitBranchNm()).Raw
		_, isOk := data.List[gitBranch]

		if !isOk {
			data.List[gitBranch] = 1
		} else {
			data.List[gitBranch] += 1
		}

		gitTag := fmt.Sprintf("%s-%d", gitBranch, data.List[gitBranch])

		fnPanic.HasError(exec.Command("git", "commit", "--all", "-m", fmt.Sprintf("%s\n%s version up", gitTag, commitMsg)).Run())
		fnPanic.HasError(exec.Command("git", "push").Run())
		fnPanic.HasError(exec.Command("git", "tag", gitTag).Run())
		fnPanic.HasError(exec.Command("git", "push", "origin", gitTag).Run())

		fnPanic.HasError(mng.Save(data))
	}

	return
}
