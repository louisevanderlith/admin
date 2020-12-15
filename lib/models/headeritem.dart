// import 'dart:html';

// import 'package:mango_ui/bodies/header.dart';
// import 'package:mango_ui/bodies/key.dart';
// import 'package:mango_ui/services/uploadapi.dart';

// class HeaderItem {
//   FileUploadInputElement _image;
//   TextInputElement _heading;
//   TextAreaElement _text;
//   bool _loaded;

//   HeaderItem(String imageElem, String headingElem, String textElem) {
//     _image = querySelector(imageElem);
//     _heading = querySelector(headingElem);
//     _text = querySelector(textElem);

//     _loaded = _image != null && _heading != null && _text != null;

//     if (_loaded) {
//       _image.onChange.listen(uploadFile);
//     }
//   }

//   Key get imageKey {
//     return new Key(_image.dataset["id"]);
//   }

//   String get heading {
//     return _heading.value;
//   }

//   String get text {
//     return _text.value;
//   }

//   bool loaded() {
//     return _loaded;
//   }

//   Header toDTO() {
//     return new Header(heading, text, imageKey);
//   }
// }
