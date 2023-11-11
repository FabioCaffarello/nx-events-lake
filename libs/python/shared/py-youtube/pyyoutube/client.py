import io
import pytube
from pylog.log import setup_logging


logger = setup_logging(__name__, log_level="DEBUG")

class YoutubeDownloaderError(Exception):
    def __init__(self, args, kwargs):
        super().__init__(args, kwargs)


def download_to_buffer(url: str) -> bytes:
    """
    Downloads a YouTube video from the specified URL and returns it as bytes data.

    Args:
        url (str): The URL of the YouTube video to download.

    Returns:
        bytes: The downloaded video data as bytes.

    Raises:
        YoutubeDownloaderError: If there is an error during the download process.

    Note:
        This function uses the PyTube library to download videos from YouTube.
        It retrieves the highest quality MP4 video stream available and saves it in a bytes buffer.

    Example:
        >>> video_url = "https://www.youtube.com/watch?v=example_video_id"
        >>> video_data = download_to_buffer(video_url)
        >>> with open("downloaded_video.mp4", "wb") as file:
        ...     file.write(video_data)

    """
    logger.info(f"Input URL: {url}")
    youtube = pytube.YouTube(url)
    buffer = io.BytesIO()
    try:
        video_stream = (
            youtube
            .streams
            .filter(
                progressive=True,
                file_extension='mp4',
                type="video"
            )
            .order_by('resolution')
            .desc()
            .first()
        )
        video_stream.stream_to_buffer(buffer)
        logger.info("Video downloaded successfully.")
    except Exception as e:
        raise YoutubeDownloaderError(f"Failed to download video: {e}")
    bytes_data = buffer.getvalue()

    buffer.close()
    return bytes_data
