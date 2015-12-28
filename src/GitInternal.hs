module GitInternal where

import System.Process
import qualified Data.Text as Text

executeCommand :: [String] -> IO String
executeCommand command = chomp <$> (readProcess process args [])
  where process = head command
        args = tail command
        chomp = Text.unpack . Text.strip . Text.pack

extractResult :: [(a, [(Int, String)])] -> String
extractResult = snd . head . snd . head
