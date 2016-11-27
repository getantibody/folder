# folder [![Build Status](https://travis-ci.org/getantibody/folder.svg?branch=master)](https://travis-ci.org/getantibody/folder)

URL to folder name parser based on antigen rules

## Examples


```golang
folder.FromURL("git@github.com:getantibody/antibody.git")
// git-AT-github.com-COLON-getantibody-SLASH-antibody.git
```

```golang
folder.ToURL("https-COLON--SLASH--SLASH-github.com-SLASH-getantibody-SLASH-folder")
// https://github.com/getantibody/folder
```
