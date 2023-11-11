# py-youtube

`py-youtube` is a Python library that allows you to easily download YouTube videos from their URLs. It uses the PyTube library to retrieve the highest quality MP4 video stream available and save it as bytes data in a buffer. This library is helpful for various use cases, such as downloading YouTube videos for offline viewing or further processing.

## Installation

You can install `py-youtube` using `nx`:

```sh
npx nx add <project> --name python-shared-py-youtube --local
```

## Usage

To download a YouTube video, you can use the `download_to_buffer` function provided by this library. Here's how to use it:

```python
from pyyoutube import download_to_buffer

# Specify the URL of the YouTube video you want to download
video_url = "https://www.youtube.com/watch?v=example_video_id"

# Download the video and get the video data as bytes
video_data = download_to_buffer(video_url)

# Save the video data to a file (e.g., downloaded_video.mp4)
with open("downloaded_video.mp4", "wb") as file:
    file.write(video_data)
```

## Exception Handling

The library defines an `YoutubeDownloaderError` exception that is raised if there is an error during the download process. You can catch and handle this exception in your code to deal with download failures gracefully.

```python
from pyyoutube import download_to_buffer, YoutubeDownloaderError

try:
    video_data = download_to_buffer(video_url)
except YoutubeDownloaderError as e:
    # Handle the download error
    print(f"Download error: {e}")
```


## Contributions

Contributions and bug reports are welcome! If you find any issues or have suggestions for improvements, please feel free to create a GitHub issue or submit a pull request.
