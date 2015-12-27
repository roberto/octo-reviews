module GitSpec where

import Test.Hspec
import Git

main :: IO ()
main = hspec $ do
  describe "currentBranch" $ do
    it "returns master" $ do
      currentBranch `shouldReturn` "master"

  describe "remoteUrl" $ do
    it "returns this repository url" $ do
      remoteUrl `shouldReturn` "git@github.com:roberto/octo-reviews.git"

  describe "owner" $ do
    it "returns this repository owner" $ do
      owner `shouldReturn` "roberto"

  describe "project" $ do
    it "returns this repository name" $ do
      project `shouldReturn` "octo-reviews"
