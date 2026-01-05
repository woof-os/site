cd src/generator
./main > ../../packages.html 

cd ../../
tailwindcss build -i ./src/tailwind.css -o ./tailwind.css
