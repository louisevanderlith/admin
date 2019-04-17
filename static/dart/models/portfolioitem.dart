import 'dart:html';

class PortfolioItem {
  FileUploadInputElement _image;
  TextInputElement _name;
  UrlInputElement _url;
  bool _loaded;

  PortfolioItem(String imageElem, String nameElem, String urlElem) {
    _image = querySelector(imageElem);
    _name = querySelector(nameElem);
    _url = querySelector(urlElem);

    _loaded = _image != null && _name != null && _url != null;
  }

  String get imageKey {
    return _image.dataset["id"];
  }

  String get name {
    return _name.value;
  }

  String get url {
    return _url.value;
  }

  bool loaded() {
    return _loaded;
  }

  Object toJson() {
    return {"imageKey": imageKey, "name": name, "url": url};
  }
}
