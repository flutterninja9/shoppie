# Shoppie

Welcome to **Shoppie** - an experimental e-commerce platform where we're exploring the power of microservice architecture. Our journey is backend-focused, with Go powering our services, and we've dabbled a bit in Flutter for a minimal frontend.

## Core Concept

- **Experiment with Microservices**: Shoppie is our playground for testing scalable and efficient service handling via microservices.
- **Go Backend**: The backbone of our experiment, utilizing Go for its performance and concurrency capabilities.
- **Flutter Frontend**: A basic frontend, more as a proof of concept rather than a fully-fledged user interface.

## Current Focus & Pending Tasks

As an experimental project, here are some aspects we're currently tinkering with:

1. **User Cart Validation**:
   - Implementing security measures to ensure that the token and userId for the user cart are aligned, reinforcing that users can access only their cart data.

2. **API Endpoint Optimization**:
   - Resolving the redundant `get-user-orders` endpoint present in both order and user services within our backend-for-frontend (BFF) layer.

3. **Admin Panel Frontend (Lightweight)**:
   - Developing a basic version of the admin panel's frontend, primarily for backend interaction and minimal admin functionalities.
