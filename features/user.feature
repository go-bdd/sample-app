Feature: User profile
    Scenario: creating a new account
        Given I create a new account with username admin
        Then the creation of the account succeeded
        When I log in to the system using username admin
        Then the logging in succeeded
        When I create a new account with username admin
        Then the creation of the account failed
