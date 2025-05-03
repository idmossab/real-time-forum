#!/bin/bash

# Colors for better output
GREEN='\033[0;32m'
RED='\033[0;31m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

BASE_URL="http://localhost:8080"

# Test user registration
test_register() {
    echo -e "${BLUE}Testing user registration...${NC}"
    echo -e "Request:"
    echo '{
        "nick_name": "testuser",
        "age": 25,
        "gender": "male",
        "first_name": "Test",
        "last_name": "User",
        "email": "test@example.com",
        "password": "password123"
    }'
    
    echo -e "\nResponse:"
    curl -s -X POST ${BASE_URL}/api/register \
        -H "Content-Type: application/json" \
        -d '{
            "nick_name": "testuser",
            "age": 25,
            "gender": "male",
            "first_name": "Test",
            "last_name": "User",
            "email": "test@example.com",
            "password": "password123"
        }'
    echo # New line after response
}

# Test user login
test_login() {
    echo -e "\n${BLUE}Testing user login...${NC}"
    echo -e "Request:"
    echo '{
        "email": "test@example.com",
        "password": "password123"
    }'
    
    echo -e "\nResponse:"
    curl -v -X POST ${BASE_URL}/api/login \
        -H "Content-Type: application/json" \
        -d '{
            "email": "test@example.com",
            "password": "password123"
        }'
    echo # New line after response
}

# Main function to run all tests
main() {
    echo -e "${GREEN}======= REAL TIME FORUM API DEBUG TESTS =======${NC}"
    
    echo -e "${BLUE}These tests show full request and response details for debugging${NC}"
    
    # Run tests
    test_register
    sleep 2
    test_login
    
    echo -e "\n${GREEN}======= TESTS COMPLETED =======${NC}"
}

# Run the main function
main