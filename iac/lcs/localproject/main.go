//revive:disable:package-comments,exported
package main

import (
	"github.com/pulumi/pulumi-github/sdk/v5/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {

	pulumi.Run(func(ctx *pulumi.Context) error {

		repositoryName := "test-mirror"
		repository, err := github.NewRepository(ctx, "newRepository", &github.RepositoryArgs{
			DeleteBranchOnMerge: pulumi.Bool(true),
			Description:         pulumi.String("This is a test repository for GitLab mirroring"),
			HasIssues:           pulumi.Bool(true),
			HasProjects:         pulumi.Bool(true),
			Name:                pulumi.String(repositoryName),
			Topics:              pulumi.StringArray{pulumi.String("pulumi"), pulumi.String("dagger"), pulumi.String("github"), pulumi.String("gitlab"), pulumi.String("test")},
			Visibility:          pulumi.String("public"),
		})
		if err != nil {
			return err
		}

		_, err = github.NewBranchProtection(ctx, "branchProtection", &github.BranchProtectionArgs{
			RepositoryId:          repository.NodeId,
			Pattern:               pulumi.String("main"),
			RequiredLinearHistory: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}

		_, err = github.NewIssueLabel(ctx, "newIssueLabelGhActions", &github.IssueLabelArgs{
			Color:       pulumi.String("E66E01"),
			Description: pulumi.String("This issue is related to github-actions dependencies"),
			Name:        pulumi.String("github-actions dependencies"),
			Repository:  repository.Name,
		})
		if err != nil {
			return err
		}

		_, err = github.NewIssueLabel(ctx, "newIssueLabelGoModules", &github.IssueLabelArgs{
			Color:       pulumi.String("9BE688"),
			Description: pulumi.String("This issue is related to go modules dependencies"),
			Name:        pulumi.String("go-modules dependencies"),
			Repository:  repository.Name,
		})
		if err != nil {
			return err
		}

		// _, err = github.GetActionsPublicKey(ctx, &github.GetActionsPublicKeyArgs{
		// 	Repository: repositoryName,
		// }, nil)
		// if err != nil {
		// 	return err
		// }

		// _, err = github.NewActionsSecret(ctx, "newActionSecret", &github.ActionsSecretArgs{
		// 	Repository: pulumi.String(repositoryName),
		// 	SecretName: pulumi.String("TOKEN"),
		// })
		// if err != nil {
		// 	return err
		// }

		ctx.Export("repository", repository.Name)
		ctx.Export("repositoryUrl", repository.HtmlUrl)

		return nil
	})

}
