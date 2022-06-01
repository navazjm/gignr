<!-- PROJECT LOGO -->
<br />
<p align="center">
  <img src="https://michaelnavs-readme.s3.us-east-2.amazonaws.com/gignr.jpg" alt="Logo" width="500">

  <h3 align="center">gignr (Gee Ignore)</h3>

  <p align="center">
  Generate .gitignore files from your terminal
  </p>
</p>

## About gignr

In the past, I used [toptal/gitignore.io](https://github.com/toptal/gitignore.io) to generate a .gitignore file for all my projects. However,
they decided to disable their tool if you use an ad blocker. In my opinion, this is the wrong approach
for a free tool, so I am creating a free and open source replacement tool. For now, the plan is to create
a simple, usable cli tool to generate a .gitignore file. I would love to implement new features as long
as the mvp for gignr stays the same. The goal for gignr is as follows:

- Easy, simple to use
- Fast, no waiting as the user is trying to create their next best idea.

## Installation

## Usage

- List all available templates to use to generate a .gitignore file for your project

```sh 
gignr ls
```

- Generate a new .gitinore file with a single template, use the -t or --template flag

```sh
gignr genrate -t <TEMPLATE_NAME>
```

- Generate a new .gitinore file with multiple templates

```sh 
gignr genrate -t <TEMPLATE_NAME> -t <TEMPLATE_NAME> ...
```

- Append a template to an existing .gitinore file, use the -a or --append flag

```sh 
gignr genrate -a -t <TEMPLATE_NAME>
```

- Append multiple templates to an existing .gitinore file

```sh 
gignr genrate -a -t <TEMPLATE_NAME> -t <TEMPLATE_NAME> ...
```

## Contribute

All contributions are welcome! Just open a pull request. Please read [CONTRIBUTING.md](./CONTRIBUTING.md)

## License

Distributed under the MIT License. See [LICENSE](./LICENSE) for more information.

## Contact

Michael Navarro - [@michaeljnavs](https://twitter.com/michaeljnavs) - michaelnavs@gmail.com

Project Link: [https://github.com/michaelnavs/gignr](https://github.com/michaelnavs/gignr)
