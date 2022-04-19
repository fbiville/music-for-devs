sine-generator
==============

Sends a "pure" sine wave to the system's default audio output.

Compiling
---------

This project requires Java 17 and a recent version of Maven (3.8+).

```
mvn clean package
```

Usage
-----

```
java -jar target/sine-generator-*.jar [frequency in Hz] [duration in ms] [volume]
```

- `frequency` is a floating point value in Hz
- `duration` is an integer value of milliseconds
- `volume` is an integer between 0 and 256
