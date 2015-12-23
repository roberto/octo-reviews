module Git where

import GitInternal
import System.Process

currentBranch :: IO String
currentBranch = executeCommand ["git", "rev-parse", "--abbrev-ref", "HEAD"]

remoteUrl :: IO String
remoteUrl = executeCommand ["git", "config", "--get", "remote.origin.url"]
