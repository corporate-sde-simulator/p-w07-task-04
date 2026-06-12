# Beginner Explanatory Guide: PLATFORM-2958: Build blue-green deployment switcher

> **Task Type**: Product Task  
> **Domain/Focus**: Backend Development, Deployment Strategies

---

## 1. The Goal (In-Depth Beginner Explanation)

### The Core Problem
In modern software development, ensuring that applications are always available and running smoothly is crucial. The task at hand involves creating a `DeploymentSwitcher` that facilitates blue-green deployments. Currently, the system lacks a mechanism to seamlessly switch between two environments (blue and green) without downtime. This is problematic because if a new version of the application is deployed and it contains bugs, users may experience errors or downtime, leading to a poor user experience and potential loss of revenue.

The `HealthGate` module is designed to check the health of services, but without the `DeploymentSwitcher`, there is no automated way to route traffic to the healthy environment after deployment. This means that if a deployment fails, the system cannot automatically revert to the previous stable version, which can lead to prolonged outages. Fixing this issue is vital for maintaining high availability and reliability of the application, ensuring that users always have access to a functioning version of the service.

### Jargon Buster (Key Terms Explained)
* **Blue-Green Deployment**: This is a deployment strategy that reduces downtime and risk by running two identical production environments, called "blue" and "green." At any time, one environment is live (serving users), while the other is idle. When a new version is ready, it is deployed to the idle environment, and traffic is switched over only after successful health checks.
  
* **Health Checks**: These are automated tests that verify whether a service is functioning correctly. They can check various aspects, such as whether the service is responding to requests, whether it can connect to a database, and whether it is performing within expected parameters. If a health check fails, it indicates that the service is not ready to handle user traffic.

* **Rollback**: This is the process of reverting to a previous version of the application after a deployment fails. It is crucial for maintaining system stability, as it allows developers to quickly restore service to a known good state without significant downtime.

* **Deployment History**: This refers to the record of all deployments made to the application, including timestamps, versions, and the status of each deployment (successful or failed). Keeping track of deployment history is important for auditing and troubleshooting.

### Expected Outcome
After implementing the `DeploymentSwitcher`, the system should be able to:
- Deploy new versions of the application to the inactive environment (either blue or green).
- Automatically switch traffic to the newly deployed environment only after all health checks pass, ensuring that users are not directed to a faulty version.
- Allow for easy rollback to the previous environment if the new deployment fails or if issues are detected.
- Maintain a log of deployment history for future reference.

**Before vs. After**:
- **Before**: Deployments could lead to downtime or errors if the new version is faulty, and there is no automated way to revert to the previous version.
- **After**: Deployments are seamless, with automatic health checks ensuring that only stable versions are live, and rollbacks can be performed quickly if needed.

---

## 2. Related Coding Concepts & Syntax (50% Theory, 50% Practice)

### Concept 1: Functions in Go
#### 📘 Theoretical Overview (50%)
* **Why it exists**: Functions are fundamental building blocks in programming that allow us to encapsulate code for reuse. They help organize code into manageable sections, making it easier to read, maintain, and debug. Without functions, code would become repetitive and harder to manage, leading to increased errors and complexity.
* **Key Mechanisms**: In Go, functions are defined using the `func` keyword, followed by the function name, parameters, and return types. They can take inputs (parameters) and return outputs (return values), allowing for flexible and modular code design.

#### 💻 Syntax & Practical Examples (50%)
* **Language Syntax**:
  ```go
  func functionName(parameter1 type, parameter2 type) returnType {
      // Function body
      return value
  }
  ```
  - `func`: This keyword indicates the start of a function definition.
  - `functionName`: The name of the function, which should be descriptive of its purpose.
  - `parameter1 type`: The input to the function, where `parameter1` is the name and `type` is the data type (e.g., `int`, `string`).
  - `returnType`: The type of value the function will return.

* **Real-World Application**:
  ```go
  func add(a int, b int) int {
      return a + b
  }
  ```
  In this example, the `add` function takes two integers as input and returns their sum. This function can be reused anywhere in the code where addition is needed.

---

## 3. Step-by-Step Logic & Walkthrough

1. **Step 1: Locate and Analyze the Target File**
   * Navigate to the `p-w07-task-04` folder and open the `deploymentSwitcher.go` file. This is where you will implement the functionality for the `DeploymentSwitcher`.
   * Look for the sections marked with `TODO` comments, as these indicate where you need to add your code.

2. **Step 2: Input Verification & Validation**
   * Before implementing the core logic, ensure that you handle edge cases. For example, check if the environments (blue and green) are properly initialized and if the health checks are defined.

3. **Step 3: Core Implementation / Modification**
   * Implement the `Deploy()` function to deploy to the inactive environment. This function should take the version of the application as input and update the inactive environment.
   * Implement the `Switch()` function to route traffic to the new environment only after confirming that all health checks pass. This will involve calling the `HealthGate` module.
   * Implement the `Rollback()` function to revert to the previous environment if the new deployment fails.

4. **Step 4: Output Verification & Testing**
   * After implementing the functions, run the unit tests provided in the repository to ensure that all tests pass. This will verify that your implementation works as expected and meets the acceptance criteria.

---

## 4. Detailed Walkthrough of Test Cases

### Test Case 1: Standard / Success Case
* **Description**: This test checks if the `Deploy()` function successfully deploys a new version to the inactive environment.
* **Inputs**:
  ```json
  {
      "version": "1.0.1",
      "environment": "green"
  }
  ```
* **Step-by-Step Execution Trace**:
  1. The `Deploy()` function is called with version `1.0.1` and the environment `green`.
  2. The function checks if the `green` environment is inactive.
  3. The deployment process begins, and the new version is installed in the `green` environment.
  4. The function logs the deployment in the deployment history.
* **Expected Output**: The function returns a success message indicating that the deployment was successful.

### Test Case 2: Edge Case / Validation Fail
* **Description**: This test checks how the system handles a deployment attempt when the environment is already active.
* **Inputs**:
  ```json
  {
      "version": "1.0.2",
      "environment": "blue"
  }
  ```
* **Step-by-Step Execution Trace**:
  1. The `Deploy()` function is called with version `1.0.2` and the environment `blue`.
  2. The function checks if the `blue` environment is inactive.
  3. Since the `blue` environment is active, the function raises an error indicating that deployment cannot proceed.
  4. The execution is halted, and no changes are made.
* **Expected Output**: The function returns an error message stating that the deployment cannot proceed because the environment is already active.