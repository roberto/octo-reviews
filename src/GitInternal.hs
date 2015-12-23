module GitInternal where

import System.Process

chomp :: String -> String
chomp = head . lines

executeCommand :: [String] -> IO String
executeCommand command = chomp <$> (readProcess process args [])
  where process = head command
        args = tail command

