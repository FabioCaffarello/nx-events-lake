import json
import re
from typing import Dict, List
from pylog.log import setup_logging
from bs4 import BeautifulSoup

logger = setup_logging(__name__)


def _get_soup(html: str) -> BeautifulSoup:
    """
    Get the BeautifulSoup object for the HTML string.

    Args:
        html (str): The HTML string.

    Returns:
        BeautifulSoup: The BeautifulSoup object.

    """
    try:
        soup = BeautifulSoup(html, "html.parser", from_encoding="utf-8")
        return soup
    except Exception as e:
        logger.error(f"Error getting soup: {e}")
        return None

def _find_all_by_tag_within_parent(soup: BeautifulSoup, parent_tag: str, child_tag: str) -> List[str]:
    """
    Find all child tags within a specific parent tag in the HTML.

    Args:
        parent_tag (str): The parent tag to find.
        child_tag (str): The child tag to find within the parent.

    Returns:
        List[str]: The list of child tags within the specified parent tag.

    """
    try:
        parent_tags = soup.find_all(parent_tag)
        child_tags = [parent.find(child_tag) for parent in parent_tags]
        return child_tags
    except Exception as e:
        logger.error(f"Error finding child tags within parent: {e}")
        return None


def get_href_data_from_html(html: str) -> Dict[str, str]:
    """
    Extract all href content from the <a> tags under <li> tags in the given HTML.

    Args:
        html (str): The HTML content.

    Returns:
        List[str]: List of href content.

    """
    logger.info("Getting href data from HTML")
    soup = _get_soup(html)

    # Find all "a" tags within "li" tags in the HTML
    li_a_tags = _find_all_by_tag_within_parent(soup, "li", "a")

    # Extract href content from each "a" tag
    hrefs = [a.get("href") for a in li_a_tags if a is not None and a.get("href") is not None]
    hrefs = [href for href in hrefs if "/Company" in href]

    return hrefs


def _get_pdf_target_year(soup: BeautifulSoup) -> str:
    pdf_span = soup.find("span", text="View PDF")

    if pdf_span:
        # Navigate to the parent <a> tag
        pdf_a_tag = pdf_span.find_parent("a")

        if pdf_a_tag and "onclick" in pdf_a_tag.attrs:
            # Extract the year using a regular expression
            match = re.search(r"open_rating\(\d+, '.* (\d+)", pdf_a_tag["onclick"])
            if match:
                year = match.group(1)
                logger.info(f"year: {year}")
                return year
            else:
                logger.warning("Year not found in onclick attribute")


def get_document_download_target(html: str) -> str:
    logger.info("Getting document download target")
    soup = _get_soup(html)
    # Find the first <span> element with class "btn_archived view_annual_report"
    span = soup.find('span', class_='btn_archived view_annual_report')
    if span is None:
        return None

    href = span.find('a')['href']
    match = re.search(r'[a-zA-Z_]+(?=\d*\.)', href)

    if match is None:
        return None

    pdf_target_year = _get_pdf_target_year(soup)
    if pdf_target_year is None:
        return None

    target_file = f"{match.group()}{pdf_target_year}.pdf"
    logger.info(f"target_file: {target_file}")
    return target_file

# def get_document_download_target(html: str) -> str:
#     logger.info("Getting document download target")
#     soup = _get_soup(html)
#     script_target = soup.find("script", type=re.compile("application/ld\+json"))
#     if script_target is None:
#         return None
#     script_text = script_target.string
#     try:
#         script_json = json.loads(script_text)
#     except json.JSONDecodeError as err:
#         logger.error(f"Error parsing JSON: {err}")
#         return None


#     # Find the first <span> element with class "btn_archived view_annual_report"
#     span = soup.find('span', class_='btn_archived view_annual_report')
#     if span is None:
#         return None

#     href = span.find('a')['href']


#     content_url = script_json["logo"]["contentUrl"].split("/")[-1].split(".")[0]
#     logger.info(f"content_url: {content_url}")

#     pdf_target_year = _get_pdf_target_year(soup)
#     if pdf_target_year is None:
#         return None

#     target_file = f"{content_url}_{pdf_target_year}.pdf"
#     logger.info(f"target_file: {target_file}")
#     return target_file
