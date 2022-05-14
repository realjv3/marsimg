# Returns JSON of most recent 10 days worth of photos from Mars rovers

Rover *name* and *camera* args can be set when executing the binary, e.g `./marsimg curiosity navcam`

The *camera* argument can be omitted, or both arguments can be omitted.

The default arg for rover is `curiosity`. Currently this is the only supported rover.
* `fhaz` - Front Hazard Avoidance Camera
* `rhaz` - Rear Hazard Avoidance Camera
* `mast` - Mast Camera
* `navcam` - Navigation Camera

 For frequent use, set the `apiKey` var in the `main.go` file with your https://api.nasa.gov API key before compiling.