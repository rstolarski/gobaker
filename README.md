# Gobaker

Gobaker is a Go program thats intented to solve 

Standard 3D modeling programs (e.g. Blender or 3DS Max) can do baking process of highpoly mesh to a lowpoly model, but they lack alpha checking.

When you want to bake 2 single faces with alpha texture on them, each program will write whatever is on the fartest triangle from the mesh, ignoring each triangle underneath. This program was made to solve this issue.

![Issue with expected result](https://i.imgur.com/SN0Ds6H.png)

Repository contains two programs, one written in Go which represent backed for baking operation and second one which is GUI app written is C#.

## Requirements
* Each texture has to be either in .png or .jpg format
* Each mesh has to be triangulated

**Currently Mask/ID map blue channel is multiplied by vertex alpha value and the Mask/ID map alpha channel is the depth map, it is not rendered from highpoly texture**

## Basic usage
This application needs following files:
* Lowpoly mesh in OBJ file format
* Highpoly mesh in OBJ file format
* Highpoly mesh in PLY file format with saved vertex alpha color
* Textures for lowpoly mesh:
    * Albedo/Diffuse with Opacity in alpha channel
    * Mask/Id map

## Console application
For console application specifi flags are supported.
Keyboard controls are indicated below.

### Required flags
Description | Flag
--- | :---:
Render size of the output image | _-s_
Path to lowpoly mesh            | _-l_
Path to highpoly mesh           | _-h_
Path to PLY mesh                | _-hp_
Rendered image output directory | _-o_

### Optional flags
Description | Flag
--- | :---:
Turn on cpu profiling | _-cpuP_
Turn on memory profiling | _-memP_
Turn on trace profiling | _-traceP_
Use half of available CPU cores. Otherwise use all use can | _-useHalfCPU_

## GUI application
For GUI version of this app setting every flag from above (except optional flags) can be set using buttons. This version use all available CPU cores.

![GUI applcation](https://i.imgur.com/YzWjlf6.png)
