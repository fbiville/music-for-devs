# A final example...

sample :drum_tom_hi_hard
sleep 0.125
sample :drum_tom_mid_hard
sleep 0.25
sample :drum_tom_hi_hard
sleep 0.125
sample :drum_tom_hi_hard
sleep 0.125
sample :drum_tom_mid_hard
sleep 0.125
sample :drum_tom_mid_hard
sleep 0.125
sample :drum_tom_low_hard
sleep 0.125
sample :drum_tom_mid_hard
sleep 0.125
sample :drum_tom_low_hard
sleep 0.125

in_thread do
  loop do
    sample :drum_bass_hard
    sample :drum_bass_soft
    sleep 0.5
    sample :drum_snare_soft
    sample :drum_snare_hard
    sleep 0.5
  end
end
in_thread do
  loop do
    sleep 0.25
    sample :drum_cymbal_closed
    sleep 0.5
    sample :drum_cymbal_closed
    sleep 0.25
  end
end

in_thread do
  use_synth :fm
  loop do
    play chord(:Bb5,:M)
    sleep 0.75
    play chord(:C6,:M)
    sleep 1.25
    play chord(:C6,:M)
    sleep 0.75
    play chord(:D6,:m)
    sleep 1.25
    play chord(:Bb5,:M)
    sleep 0.75
    play chord(:C6,:M)
    sleep 1.25
    play chord(:A5,:m)
    sleep 0.75
    play chord(:Bb5,:M)
    sleep 1.25
  end
end

in_thread do
  with_synth :fm  do
    loop do
      play(:Bb2)
      sleep 0.5
      play(:Bb2)
      sleep 0.25
      play(:Bb2)
      sleep 0.5
      play(:G3)
      sleep 0.25
      play(:F3)
      sleep 0.5
      play(:E3)
      sleep 0.5
      play(:E3)
      sleep 1
      play(:C3)
      sleep 0.25
      play(:E3)
      sleep 0.25
      
      play(:C3)
      sleep 0.5
      play(:C3)
      sleep 0.25
      play(:C2)
      sleep 0.5
      play(:G3)
      sleep 0.25
      play(:F3)
      sleep 1
      play(:F3)
      sleep 0.25
      play(:F3)
      sleep 0.25
      play(:F3)
      sleep 0.25
      play(:G3)
      sleep 0.5
      play(:E3)
      sleep 0.25
    end
  end
end


in_thread do
  with_synth :hoover do
    with_fx :distortion, distort:0 do
      sleep 8
      loop do
        play :F7
        sleep 0.75
        play :G7
        sleep 1.25
        play :G7
        sleep 0.75
        play :A7
        sleep 0.75
        play :C8, sustain:0.1
        sleep 0.125
        play :Bb7, sustain:0.1
        sleep 0.125
        play :A7, sustain:0.2
        sleep 0.25
        
        play :F7
        sleep 0.75
        play :G7
        sleep 1.25
        play :E7
        sleep 0.75
        play :F7
        sleep 1.25
      end
    end
  end
end

