import 'dart:html';
import 'package:mango_ui/bodies/key.dart';
import 'package:mango_ui/bodies/portfolio.dart';
import 'package:mango_ui/services/uploadapi.dart';

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

  Key get imageKey {
    return new Key(_image.dataset["id"]);
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

  Portfolio toDTO() {
    return new Portfolio(imageKey, url, name);
  }
}
