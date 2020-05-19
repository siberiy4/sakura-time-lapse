#https://qiita.com/riocampos/items/2f4fe927b5cf99aff767

#ffmpeg -f image2 -r 20 -i  source%04d.jpg -r 40 -an -vcodec libx264  -pix_fmt yuv420p video.mp4

./ffmpeg-4.2.2-amd64-static/ffmpeg -f image2 -r 20 -i  test/source%04d.jpg -r 40 -an -vcodec libx264  -pix_fmt yuv420p video.mp4