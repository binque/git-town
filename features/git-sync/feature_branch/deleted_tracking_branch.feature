Feature: git-sync restores deleted tracking branch

  Scenario: without a remote branch
    Given I have a feature branch named "feature"
    And the following commits exist in my repository
      | branch  | location         | message        | file name    |
      | feature | local and remote | feature commit | feature_file |
    And the "feature" branch gets deleted on the remote
    And I am on the "feature" branch
    When I run `git sync`
    Then I am still on the "feature" branch
    And I have the following commits
      | branch  | location         | message        | files        |
      | feature | local and remote | feature commit | feature_file |