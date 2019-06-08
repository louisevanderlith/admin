import 'dart:html';
import '../services/uploadapi.dart';

class PortfolioItem {
  FileUploadInputElement _image;
  TextInputElement _name;
  UrlInputElement _url;
  bool _loaded;

  PortfolioItem(
      String imageElem, String nameElem, String urlElem) {
    _image = querySelector(imageElem);
    _name = querySelector(nameElem);
    _url = querySelector(urlElem);

    _loaded = _image != null && _name != null && _url != null;

    if (_loaded) {
      _image.onChange.listen(uploadFile);
    }
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
    return {"ImageKey": imageKey, "Name": name, "URL": url};
  }
}
