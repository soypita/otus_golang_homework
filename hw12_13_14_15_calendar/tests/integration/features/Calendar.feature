Feature: Calendar service
  As a client of calendar service
  I want to work with calendar events through gRPC API

  Scenario: Check health endpoint
    When I send "GET" request to "http://localhost:8080/health"
    Then Response code should be 200
    And Status in response should be "UP"

  Scenario: Create event
    When I call createEvent method
    Then Response should have event ID

  Scenario: Create event duplicate error
    Given there is event with date "2020-06-28T00:40:08.000Z"
    When I call createEvent method for date "2020-06-28T00:40:08.000Z"
    Then I get error response

  Scenario: Update event
    Given there is event with date "2020-06-28T00:40:08.000Z"
    When I call updateEvent method
    Then I get success response

  Scenario: Update event with error for unknown event id
    When I call updateEvent method for "f47ac10b-58cc-0372-8567-0e02b2c3d479"
    Then I get error response

  Scenario: Delete event
    Given there is event with date "2020-06-28T00:40:08.000Z"
    When I call deleteEvent method
    Then I get success response

  Scenario: Delete event for unknown event id
    When I call deleteEvent method for "f47ac10b-58cc-0372-8567-0e02b2c3d479"
    Then I get success response

  Scenario: Find all day events
    Given there is event with date "2020-06-28T10:40:08.000Z"
    When I call findDayEvents method for day "2020-06-28T00:00:00.000Z"
    Then I get success response
    And Events response size should be 1

  Scenario: Find all day events empty response
    Given there is event with date "2020-06-28T10:40:08.000Z"
    When I call findDayEvents method for day "2020-05-28T00:00:00.000Z"
    Then I get success response
    And Events response size should be 0

  Scenario: Find week events
    Given there is event with date "2020-06-28T10:40:08.000Z"
    When I call findWeekEvents method for day "2020-06-25T00:00:00.000Z"
    Then I get success response
    And Events response size should be 1

  Scenario: Find week events empty response
    Given there is event with date "2020-06-28T10:40:08.000Z"
    When I call findWeekEvents method for day "2020-06-20T00:00:00.000Z"
    Then I get success response
    And Events response size should be 0

  Scenario: Find month events
    Given there is event with date "2020-06-28T10:40:08.000Z"
    When I call findMonthEvents method for day "2020-05-30T00:00:00.000Z"
    Then I get success response
    And Events response size should be 1

  Scenario: Find month events empty response
    Given there is event with date "2020-06-28T10:40:08.000Z"
    When I call findMonthEvents method for day "2019-06-30T00:00:00.000Z"
    Then I get success response
    And Events response size should be 0

  Scenario: Receive event day notification
    When I call createEvent method for current date
    Then Response should have event ID
    And I should receive event notification

  Scenario: Clear year old notification
    When I call createEvent method for old date
    Then Response should have event ID
    And Events should clear async