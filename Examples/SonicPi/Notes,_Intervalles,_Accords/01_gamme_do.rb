#
# Montons la gamme de Do Majeur...
#

define :demi_temps do
  sleep 0.5
end

use_synth :sine
# use_synth :saw
# use_synth :square

play :C3    #----------------------------------------
demi_temps  #   /\
play :D3    #---||-----------------------------------
demi_temps  #   |/                               C4
play :E3    #---/----------------------------B3------
demi_temps  #  /|                        A3
play :F3    #--||/\------------------G3--------------
demi_temps  #  || |              F3
play :G3    #--\|-/----------E3----------------------
demi_temps  #   \/       D3
play :A3    #   |   -C3-
demi_temps  #  \/
play :B3    #
demi_temps
play :C4
