module Git where

import GitInternal
import Text.RegexPR

currentBranch :: IO String
currentBranch = executeCommand ["git", "rev-parse", "--abbrev-ref", "HEAD"]

remoteUrl :: IO String
remoteUrl = executeCommand ["git", "config", "--get", "remote.origin.url"]

owner :: IO String
owner = extractResult <$> matchOwner <$> remoteUrl where
        matchOwner = gmatchRegexPR ":(.*)/"

project :: IO String
project = extractResult <$> matchProject <$> remoteUrl where
          matchProject = gmatchRegexPR ":.*/(.*).git"
