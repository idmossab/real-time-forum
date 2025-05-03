#!/bin/bash

# Test user registration API
echo "Testing user registration..."
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "nick_name": "testuser",
    "age": 25,
    "gender": "male",
    "first_name": "Test",
    "last_name": "User",
    "email": "test@example.com",
    "password": "password123"
  }' \
  http://localhost:8080/api/register
echo

# Run with: chmod +x test-api.sh && ./test-api.sh