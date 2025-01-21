# OpenComply Google Workspace Integration SOP

This Standard Operating Procedure (SOP) outlines the steps to integrate your Google Workspace with OpenComply by creating a Service Account with read-only access. The process includes creating the service account, selecting the owner's email for administration impersonation, generating keys, enabling read-only access, and configuring OpenComply with the necessary credentials.

## Steps

### 1. Create Service Account

1. **Open Google Cloud Console**
   - Visit [Google Cloud Console](https://console.cloud.google.com/) and select your project.

2. **Navigate to Service Accounts**
   - Go to `IAM & Admin` > `Service Accounts`.

3. **Create Account**
   - Click **Create Service Account**.
   - Enter a **Name** and **Description**.
   - Click **Create**.

### 2. Assign Admin Access and Generate Keys

1. **Assign Roles**
   - Skip assigning roles at this stage to maintain a lean setup.

2. **Create Key**
   - In the Service Accounts list, locate your newly created account.
   - Click on the service account, then go to the **Keys** tab.
   - Click **Add Key** > **Create New Key**.
   - Choose **JSON** and click **Create** to download the key file.
   - **Store the key file securely**.

3. **Handle Policy Restrictions (If Needed)**
   - If you encounter issues creating the service account or keys, adjust the organization policies:
     1. **Access Organization Policies**
        - Navigate to `IAM & Admin` > `Organization Policies` in Google Cloud Console.
     2. **Modify Policies**
        - Locate **Disable service account creation** and **Disable service account key creation**.
        - Click **Override** for each policy at the project level.
        - Set both policies to **Not Enforced**.
        - Click **Save**.
     3. **Verify Changes**
        - Retry creating the service account and generating the key.

### 3. Add API Access to Workspaces

1. **Enable Domain-Wide Delegation**
   - In the Service Account details, navigate to the **Details** tab.
   - Click **Enable G Suite Domain-wide Delegation**.
   - Provide a **Product Name** and save.

2. **Authorize API Access**
   - Go to [admin.google.com](https://admin.google.com/) and log in with admin credentials.
   - Navigate to `Security` > `API Controls` > `Domain-wide Delegation`.
   - Click **Add New**.
   - Enter the **Client ID** from the service account.
   - Add the required **Scopes**:
     ```
     https://www.googleapis.com/auth/admin.directory.user
     https://www.googleapis.com/auth/admin.directory.group
     https://www.googleapis.com/auth/admin.directory.user.readonly
     https://www.googleapis.com/auth/admin.directory.group.readonly
     https://www.googleapis.com/auth/directory.readonly
     https://www.googleapis.com/auth/contacts.readonly
     https://www.googleapis.com/auth/admin.directory.orgunit.readonly
     https://www.googleapis.com/auth/admin.directory.domain.readonly
     https://www.googleapis.com/auth/admin.directory.device.chromeos.readonly
     https://www.googleapis.com/auth/admin.directory.device.mobile.readonly
     https://www.googleapis.com/auth/admin.directory.customer.readonly
     https://www.googleapis.com/auth/admin.directory.rolemanagement.readonly
     ```
   - Click **Authorize** to grant API access.

### 4. Add Credentials to OpenComply

1. **Access OpenComply Dashboard**
   - Open your web browser and navigate to the OpenComply portal.
   - Log in with your administrator credentials.

2. **Navigate to Integrations**
   - Go to `Integrations` > `Google Workspace`.

3. **Enter Credentials**
   - **Key File**: Upload the downloaded JSON key file.
   - **Admin Email**: Enter the **owner's email** for administration impersonation (e.g., `owner@example.com`).
   - **Customer ID**:
     - In Google Admin Console, go to `Account` > `Account Settings` > `Profile`.
     - Note the **Customer ID**.

4. **Save Integration**
   - Click **Save** to complete the integration setup.

## Conclusion

You have successfully integrated your Google Workspace with OpenComply by creating a Service Account with read-only access. This setup enables OpenComply to manage and provide visibility over your Google Workspace resources using the owner's email for administration impersonation.