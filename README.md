# GTD

## Requirements

0. inbox support
1. add required time tag
2. add required energy tag
3. add custom tags (context)
4. auto synchronization
5. "waiting for" list
6. "someday/maybe" list
7. 'favourite' toggle
8. there has to be the single next list
9. add projects
10. add steps into the project
11. break steps in the project into *next* type tasks

## API

Inbox:
- `/inbox`: GET list of all tasks
- `/inbox`: POST save task
- `/inbox/{id}`: DELETE task with id

Clarify:
- `/task`: GET list of next tasks
- `/task`: POST a new next task
- `/task`: PUT update a next task

Projects:
- `/project`: GET list of projects
- `/project`: POST create a new project
- `/project/{id}`: DELETE a project
- `/project/step`: POST create a step for a project

Tags:
- `/tag`: GET get a list of tags

Box:
- `/box`: GET a list of boxes
- `/box/{id}`: get all tasks from the box

## Architecture

```plantuml
class InboxItem {
    - id: Long
    - message: String
}

enum Energy {
    LOW, MID, HIGH
}

class Tag {
    - id: Long
    - name: String
}

class ProjectStep {
    - id: Long
    - tasks: List<Task>
}

ProjectStep "use" --> Task

class Project {
    - id: Long
    - name: String
    - steps: List<ProjectStep>
}

Project "use" --> ProjectStep

enum BoxType {
    NEXT, WAITING, SOMEDAY_MAYBE
}

class Box {
    - id: Long
    - type: BoxType
}

Box "use" --> BoxType

class Task {
    - id: Long
    - project: Project?
    - box: Box?
    - message: String
    - time: Long             // time in milliseconds
    - energy: Energy
    - custom_tags: List<Tag>
}

Task "use" --> Energy
Task "use" --> Tag
Task "use" --> Project
Task "use" --> Box
```
