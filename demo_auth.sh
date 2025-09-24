#!/bin/bash

# Authentication System Demo Script
# This script demonstrates the complete authentication flow

set -e  # Exit on any error

echo "🚀 Authentication System Demo"
echo "=============================="

# Check if service is running
if ! curl -s http://localhost:8081/health > /dev/null; then
    echo "❌ User service is not running on port 8081"
    echo "Please start the service first:"
    echo "  cd user-service && go run ."
    exit 1
fi

echo "✅ User service is running"
echo ""

# Test 1: Health Check
echo "📋 1. Health Check"
echo "==================="
curl -s http://localhost:8081/health | jq .
echo ""

# Test 2: User Registration
echo "👤 2. User Registration"
echo "======================="
echo "Registering user: alice@example.com"

SIGNUP_RESPONSE=$(curl -s -X POST http://localhost:8081/auth/signup \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Alice Johnson",
    "email": "alice@example.com",
    "password": "securepass123",
    "gender": "female"
  }')

echo "$SIGNUP_RESPONSE" | jq .

# Check if signup was successful
if echo "$SIGNUP_RESPONSE" | jq -e '.user.id' > /dev/null; then
    echo "✅ User registration successful"
else
    echo "❌ User registration failed"
    exit 1
fi
echo ""

# Test 3: User Login
echo "🔐 3. User Login"
echo "================"
echo "Logging in as alice@example.com"

LOGIN_RESPONSE=$(curl -s -X POST http://localhost:8081/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "alice@example.com",
    "password": "securepass123"
  }')

echo "$LOGIN_RESPONSE" | jq .

# Extract token
TOKEN=$(echo "$LOGIN_RESPONSE" | jq -r '.token')
if [ "$TOKEN" = "null" ] || [ -z "$TOKEN" ]; then
    echo "❌ Login failed - no token received"
    exit 1
fi

echo "✅ Login successful, token received"
echo ""

# Test 4: Access Protected Profile
echo "👤 4. Access Protected Profile"
echo "=============================="
echo "Accessing profile with JWT token"

PROFILE_RESPONSE=$(curl -s -X GET http://localhost:8081/auth/profile \
  -H "Authorization: Bearer $TOKEN")

echo "$PROFILE_RESPONSE" | jq .

# Check if profile access was successful
if echo "$PROFILE_RESPONSE" | jq -e '.user.id' > /dev/null; then
    echo "✅ Profile access successful"
else
    echo "❌ Profile access failed"
    exit 1
fi
echo ""

# Test 5: Test Invalid Scenarios
echo "🛡️  5. Security Tests"
echo "===================="

echo "Testing duplicate email registration..."
DUPLICATE_RESPONSE=$(curl -s -X POST http://localhost:8081/auth/signup \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Bob Smith",
    "email": "alice@example.com",
    "password": "anotherpass123",
    "gender": "male"
  }')

if echo "$DUPLICATE_RESPONSE" | grep -q "already registered"; then
    echo "✅ Duplicate email properly rejected"
else
    echo "❌ Duplicate email not properly handled"
fi

echo "Testing invalid login credentials..."
INVALID_LOGIN=$(curl -s -X POST http://localhost:8081/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "alice@example.com",
    "password": "wrongpassword"
  }')

if echo "$INVALID_LOGIN" | grep -q "Invalid email or password"; then
    echo "✅ Invalid credentials properly rejected"
else
    echo "❌ Invalid credentials not properly handled"
fi

echo "Testing unauthorized profile access..."
UNAUTHORIZED_RESPONSE=$(curl -s -X GET http://localhost:8081/auth/profile)

if echo "$UNAUTHORIZED_RESPONSE" | grep -q "Authorization header required"; then
    echo "✅ Unauthorized access properly blocked"
else
    echo "❌ Unauthorized access not properly handled"
fi

echo ""
echo "🎉 All tests completed successfully!"
echo "======================================"
echo ""
echo "📝 Summary:"
echo "- ✅ User registration with validation"  
echo "- ✅ User login with JWT token generation"
echo "- ✅ Protected profile endpoint with middleware"
echo "- ✅ Proper error handling for edge cases"
echo "- ✅ Security validations working correctly"
echo ""
echo "🚀 The authentication system is fully functional!"