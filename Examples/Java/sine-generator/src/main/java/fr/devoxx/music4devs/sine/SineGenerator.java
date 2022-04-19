/*
 * Copyright (c) 2021 Jean-Baptiste Lièvremont
 */
package fr.devoxx.music4devs.sine;

import java.util.function.Function;
import javax.sound.sampled.AudioFormat;
import javax.sound.sampled.AudioSystem;

/**
 * <code>SineGenerator [frequency] [duration]</code>
 * Generates a monaural sine wave of the given frequency (in Hz) for the given duration (in ms).
 * Default values generate a 1 second 440Hz signal (equivalent to A<sub>4</sub> in modern standard pitch).
 */
public class SineGenerator {

  // https://en.wikipedia.org/wiki/44,100_Hz
  private static final float CDDA_SAMPLE_RATE = 44100F;

  private static final int BYTES_PER_SAMPLE = 2;
  private static final int BITS_PER_SAMPLE = 8 * BYTES_PER_SAMPLE;

  // https://en.wikipedia.org/wiki/Compact_Disc_Digital_Audio
  private static final AudioFormat CDDA_MONO = new AudioFormat(CDDA_SAMPLE_RATE, BITS_PER_SAMPLE, 1, true, false);

  public static void main(String[] args) throws Exception {
    var frequency = parseFrequency(args);
    var durationMs = parseDuration(args);
    var volume = parseVolume(args);

    var buf = new byte[BYTES_PER_SAMPLE];
    try (var line = AudioSystem.getSourceDataLine(CDDA_MONO)) {
      line.open();
      line.start();
      for (var i = 0; i < durationMs * CDDA_SAMPLE_RATE / 1000; i++) {
        var angle = i / (CDDA_SAMPLE_RATE / frequency) * 2.0 * Math.PI;
        var value = (int) (Math.sin(angle) * volume);
        buf[1] = (byte) ((value & 0xff00) >> 8);
        buf[0] = (byte) (value & 0x00ff);
        line.write(buf, 0, 2);
      }
      line.drain();
      line.stop();
    }
  }

  private static float parseFrequency(String[] args) {
    return parseArgument(args, 0, Float::valueOf, 440F);
  }

  private static int parseDuration(String[] args) {
    return parseArgument(args, 1, Integer::valueOf, 1000);
  }
  private static int parseVolume(String[] args) {
    return parseArgument(args, 2, Integer::valueOf, 256);
  }

  private static <T> T parseArgument(String[] args, int position, Function<String, T> parse, T defaultValue) {
    try {
      return parse.apply(args[position]);
    } catch(Exception e) {
      return defaultValue;
    }
  }
}
