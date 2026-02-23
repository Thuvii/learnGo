package main

import "testing"

func TestGetH1(t *testing.T) {
	html := "<html><body><h1>Test Title</h1></body></html>"
	actualOutput := getH1FromHtml(html)
	expected := "Test Title"
	if actualOutput != expected {
		t.Errorf("Expect %q, but got %q", expected, actualOutput)
	}
}

func TestGetH1Null(t *testing.T) {
	html := "<html><body><h2>Test Title</h2></body></html>"
	actualOutput := getH1FromHtml(html)
	expected := ""
	if actualOutput != expected {
		t.Errorf("Expect %q, but got %q", expected, actualOutput)
	}
}

func TestGetFirstParagraph(t *testing.T) {
	inputBody := `<html><body>
		<p>Outside paragraph.</p>
		<main>
			<p>Main paragraph.</p>
		</main>
	</body></html>`
	actual := getFirstParagraphFromHTML(inputBody)
	expected := "Main paragraph."

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetFirstParagraphNoMain(t *testing.T) {
	inputBody := `<html><body>
		<p>Outside paragraph.</p>
		<section>
			<p>Main paragraph.</p>
		</section>
	</body></html>`
	actual := getFirstParagraphFromHTML(inputBody)
	expected := "Outside paragraph."

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetFirstParagraphNull(t *testing.T) {
	inputBody := `<html><body>
		<h1>Outside paragraph.</h1>
		<section>
			<h2>Main paragraph.</h2>
		</section>
	</body></html>`
	actual := getFirstParagraphFromHTML(inputBody)
	expected := ""

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}
