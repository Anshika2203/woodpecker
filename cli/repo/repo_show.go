// Copyright 2023 Woodpecker Authors
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

package repo

import (
	"context"
	"os"
	"text/template"

	"github.com/urfave/cli/v3"

	"go.woodpecker-ci.org/woodpecker/v3/cli/common"
	"go.woodpecker-ci.org/woodpecker/v3/cli/internal"
)

var repoShowCmd = &cli.Command{
	Name:      "show",
	Usage:     "show repository information",
	ArgsUsage: "<repo-id|repo-full-name>",
	Action:    repoShow,
	Flags:     []cli.Flag{common.FormatFlag(tmplRepoInfo)},
}

func repoShow(ctx context.Context, c *cli.Command) error {
	repoIDOrFullName := c.Args().First()
	client, err := internal.NewClient(ctx, c)
	if err != nil {
		return err
	}
	repoID, err := internal.ParseRepo(client, repoIDOrFullName)
	if err != nil {
		return err
	}

	repo, err := client.Repo(repoID)
	if err != nil {
		return err
	}

	tmpl, err := template.New("_").Parse(c.String("format"))
	if err != nil {
		return err
	}
	return tmpl.Execute(os.Stdout, repo)
}

// tTemplate for repo information.
var tmplRepoInfo = `Owner: {{ .Owner }}
Repo: {{ .Name }}
URL: {{ .ForgeURL }}
Config path: {{ .Config }}
Visibility: {{ .Visibility }}
Private: {{ .IsSCMPrivate }}
Trusted: {{ .IsTrusted }}
Gated: {{ .IsGated }}
Require approval for: {{ .RequireApproval }}
Clone url: {{ .Clone }}
Allow pull-requests: {{ .AllowPullRequests }}
`
