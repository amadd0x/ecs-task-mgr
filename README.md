## Objective
Deregister and delete ephemeral task definitions.

- Discover task definition families that have > n revisions.  
- Keep the most recent revision and deregister/delete all other task definition revisions.

List all task definition families in a cluster, filtered by environment.

- I'd like to see each entry looking something like `prd-pentaho-jobs:75545`, indicating the latest revision.

I need to discover what the rate limit is for this.  I could either set this as a fixed number, or warn that we're being rate limited and adjust the time between calls dynamically if a throttling error is thrown?

