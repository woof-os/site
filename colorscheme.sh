CURRENTTHEME=$1
NEWTHEME=$2

sed -i "s/$CURRENTTHEME/$NEWTHEME/g" ~/.config/alacritty/alacritty.toml
sed -i "s/$CURRENTTHEME/$NEWTHEME/g" ~/.config/rofi/config.rasi
sed -i "s/\".*$CURRENTTHEME.*\"/\"$NEWTHEME\"/g" ~/.config/nvim/after/plugin/colors.lua

mv ~/.config/dunst/dunstrc ~/.config/dunst/dunstrc.bak
cp ~/.config/dunst/$NEWTHEME ~/.config/dunst/dunstrc

mv ~/.config/qtile/config/colors.json ~/.config/qtile/config/colors.json.bak
cp ~/.config/qtile/config/$NEWTHEME.json ~/.config/qtile/config/colors.json

[ ! -d ~/.config/starship ] && mkdir -p ~/.config/starship 
mv ~/.config/starship.toml ~/.config/starship/starship.toml.bak
cp ~/.config/starship/$NEWTHEME.toml ~/.config/starship.toml

mv ~/.config/zathura/zathurarc ~/.config/zathura/zathurarc.bak
cp ~/.config/zathura/$NEWTHEME ~/.config/zathura/zathurarc

echo "Colorscheme changed from $CURRENTTHEME to $NEWTHEME"
