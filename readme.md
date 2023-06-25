## About boxie
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/15d825c17a4c4497ba777206c18c5e3d)](https://app.codacy.com/gh/twelvee/k8sbox/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)
<img src="https://img.shields.io/docker/v/twelvee/boxie"> <br>

A tool that allows you to roll out your environments into your k8s cluster using templated specifications, monitor the activity of these services, as well as easily scrub the cluster of unused resources that you rolled out earlier.

<img src="https://i.ibb.co/5K2Bhvw/ezgif-com-crop-1.gif"><br>

## Learn boxie

All documentation is available at https://boxie.run/docs

## Contributing

Thank you for considering contributing to the boxie! The contribution guide can be found in the CONTRIBUTION.md file

### boxie <3 docker

The use of templated environments assumes integration with different ci-cd systems. That is why we offer to use a ready-made docker image for each supported version of boxie. <br>
https://hub.docker.com/r/twelvee/boxie/tags

### What boxie can already do
1. Roll out new environments
2. Update an already rolled out environment
3. Remove unused environments
4. Parsing environment variables for easy integration with any CI-CD systems
5. Show active environments
6. Describe the components of active environments

### What boxie will be able to do in the future
1. Collect statistics from your active environments
2. be more flexible for more flexible deployment
3. Automatic resource deletion by timer
4. Obtain specifications from git repositories (including private ones)
5. UI interface + REST API
..as well as many useful and easy-to-use features

## License

The boxie is open-sourced software licensed under the [MIT license](https://opensource.org/licenses/MIT).
