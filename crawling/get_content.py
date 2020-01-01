import requests
from bs4 import BeautifulSoup


def _content_to_str(s):
	a = s.find_all('p')
	if len(a) == 0:
		return [str(s.text).strip()]
	else:
		return [str(i.text).strip() for i in a]


def get_content(url):
	"""
	get content from kin.naver.com

	:param url: domain must be kin.naver.com
	:return: question and answers
			-question: title and array of lines
			-answers: array of answer that contain array of lines
	"""

	r = requests.get(url)
	if not r.ok:
		return ConnectionError

	html = r.text
	soup = BeautifulSoup(html, 'html.parser')

	question = {
		'title': _content_to_str(soup.find_all(class_='title')[0]),
		'content': _content_to_str(soup.find_all(class_='c-heading__content')[0])
	}
	answer = [_content_to_str(ans) for ans in soup.find_all(class_='_endContentsText c-heading-answer__content-user')]

	return question, answer


if __name__ == '__main__':
	a, b = get_content('https://kin.naver.com/qna/detail.nhn?d1id=13&dirId=13020103&docId=178920712')
	print(a)
	print(b)
