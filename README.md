Deregister and delete ephemeral task definitions.

ecs-task-mgr delete -t family:revision

This will deregister and delete the task definition, decrement revision by 1, and
repeat until revision 0.
