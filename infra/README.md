# Petshop infrastructure
This folder has all of the IaC for setting up the petshop application.

The idea is to create some of the missing bits in order to set up the
application; right now, there's some skeleton resources but we'll need to add
some more stuff to make it work.

# Architecture
The concept is to create an Azure Container Apps instance running our back-end
API and deploy the front-end to a static web app; if you have any suggestions on
other services or tools, now is the time to share! We're always looking for new
& innovative ways to solve problems.

# Implementation Notes
- Added a postgresql DB to handle state and data
- Password will be generated and stored in the state file
- The state file should not be local instead should be kept in an object storage
- The module could be more generalized to be able to be re-used for multiple purposes
- Output should be added for ease of access to deployment info
- Make seprate environments (dev,stage,prod) with workspaces or separate config
- Versioning to make it compatible with previous deployments
- Use a Service Principal or Managed Idenity to connect to Azure
