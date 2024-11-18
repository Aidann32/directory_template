# directory_template
A command-line tool for creating directories for various projects.

## Usage
```
dt startproject [root_directory] [project_name] [module_name]
```
**root_directory** - The directory where the project template will be created  
**project_name** - The name of the root folder of the project  
**module_name** - Name of the Go module  
### Flags
-l - The path to the layout file in json format. If the path to the layout file is not specified, the default layout will be used  

### Example of usage
```
dt startproject /home /example_project github.com/Aidann32/example
```
Path of the root directory of project will be /home/example_project  

## Default project layout
```json
{
  "config/": "config.go",
  "internal/": {
    "infra/": {
      "db/": "db.go",
      "redis/": "redis.go",
      "kafka/": "kafka.go",
      "rabbit/": "rabbit.go",
      "metrics/": "metrics.go"
    },
    "model/": "model.go",
    "handler/": "handler.go",
    "service/": "service.go",
    "repo/": "repository.go"
  },
  "pkg/": "",
  "tests/": "tests.go",
  "docs/": "docs.go",
  "main.go": "",
  "config.json": "",
  "Dockerfile": ""
}
```
