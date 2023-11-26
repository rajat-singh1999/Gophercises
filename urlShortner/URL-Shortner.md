# URL-SHORTNER

In this exercise I have created a simple URL shortner that takes input from either a YAML or a JSON file for short paths(```/path```) and URLs(```www.somwhere.com```).

The file paths are passed by the user in the command line using flags.

The ```urlshort``` folder is a package that has a file ```handler.go``` that exports the following functions used in the ```main.go``` file:

*  ```MapHandler```
*  ```YAMLHandler```

*  ```JSONHandler```


```MAPHandler``` converts data to a go map of path and urls.

```YAMLHandler``` handles case when user inputs a .yaml file.

```JSONHandler``` handles the case when user inputs a .son file

## Usage

Following are the commands to run the main.go file, note that you need to just put the name of the file not the extension:

-> Input a yaml file:

```bash
go run main.go -yaml=links
```

-> Input a json file:

```bash
go run main.go -json=links
```

-> Input both files, by default the program takes up yaml file link:

```bash
go run main.go -json=links -yaml=yamlLinks
```
