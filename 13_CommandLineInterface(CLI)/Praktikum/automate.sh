cli= "$1adhitya_ardhiansyah_$at$(date +'%a %b %d %T %Z %Y')"
mkdir "$cli"
cd "$cli"
mkdir about_me
cd about_me
mkdir personal
mkdir professional
cd ..

mkdir my_friends
mkdir my_system_info
echo "https://web.facebook.com/$2Adhitya054612"> "./about_me/personal/facebook.txt"
echo "https://www.linkedin.com/in/$3adhitya28"> "./about_me/professional/linkedin.txt"
curl "https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee4>echo "My username: $USER"> "./my_system_info/about_this_laptop.txt"
echo "With host: $(uname -a)" >> "./my_system_info/about_this_laptop.txt"
ping -c 3 google.com >"./my_system_info/internet_connection.txt"

tree