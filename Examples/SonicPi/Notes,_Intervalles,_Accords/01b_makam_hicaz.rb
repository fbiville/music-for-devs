#
# Jouons le makam Hicaz de Ré
# dans le système turc :
#  - Ré est associé à la fréquence 440 Hz, pas La !
#  - chaque ton est divisé en 9 divisions égales appelés komas (seuls le 1er, 4ème, 5ème, 8ème et 9ème sont utilisés)
#  - un octave est divisé en 53 komas
# Sonic Pi supporte la notion de "scale" (gamme), dont certaines non-occidentales !
play_pattern_timed (scale 69, :hicaz), 0.5
sleep 1

# comparons avec la gamme majeure de La (440 Hz) pour contraster
play_pattern_timed (scale 69, :major), 0.5