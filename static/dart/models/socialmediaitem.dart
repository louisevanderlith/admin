import 'dart:html';

class SocialmediaItem {
  TextInputElement _icon;
  UrlInputElement _url;
  bool _loaded;

  SocialmediaItem(String iconElem, String urlElem) {
    _icon = querySelector(iconElem);
    _url = querySelector(urlElem);

    _loaded = _icon != null && _url != null;
  }

  String get icon {
    return _icon.value;
  }

  String get url {
    return _url.value;
  }

  bool loaded() {
    return _loaded;
  }

  Object toJson() {
    return {"Icon": icon, "URL": url};
  }
}
