package git

import (
	"gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { check.TestingT(t) }

type GitSuite struct{}

var _ = check.Suite(&GitSuite{})

func (s *GitSuite) TestBranchName(c *check.C) {
	c.Assert(BranchName(), check.Equals, "master")
}

func (s *GitSuite) TestRepositoryUrl(c *check.C) {
	c.Assert(RepositoryUrl(), check.Equals, "git@github.com:roberto/octo-reviews.git")
}

func (s *GitSuite) TestRepositoryOwner(c *check.C) {
	c.Assert(RepositoryOwner(), check.Equals, "roberto")
}
