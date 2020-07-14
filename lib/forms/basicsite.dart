import 'dart:html';

import 'package:Admin.APP/models/mapitem.dart';
import 'package:mango_artifact/uploadapi.dart';
import 'package:mango_ui/keys.dart';

class BasicSiteForm {
  FormElement form;
  TextInputElement txtTitle;
  TextAreaElement txtDescription;
  FileUploadInputElement uplProfileImg;

  BasicSiteForm() {
    form = querySelector("frmBasicSite");
    txtTitle = querySelector("txtTitle");
    txtDescription = querySelector("txtDescription");
    uplProfileImg = querySelector("uplProfileImg");

    uplProfileImg.onChange.listen(uploadFile);
  }

  List<MapItem> getEndpoints() {
    var isLoaded = false;
    var result = new List<MapItem>();
    var indx = 0;

    do {
      var item =
          new MapItem('#txtEndpointName${indx}', "txtEndpointValue${indx}");

      isLoaded = item.loaded;

      if (isLoaded) {
        result.add(item);
      }

      indx++;
    } while (isLoaded);

    return result;
  }

  List<MapItem> getCodes() {
    var isLoaded = false;
    var result = new List<MapItem>();
    var indx = 0;

    do {
      var item =
      new MapItem('#txtCodeName${indx}', "txtCodeValue${indx}");

      isLoaded = item.loaded;

      if (isLoaded) {
        result.add(item);
      }

      indx++;
    } while (isLoaded);

    return result;
  }

  List<MapItem> getTerms() {var isLoaded = false;
  var result = new List<MapItem>();
  var indx = 0;

  do {
    var item =
    new MapItem('#txtTermName${indx}', "txtTermValue${indx}");

    isLoaded = item.loaded;

    if (isLoaded) {
      result.add(item);
    }

    indx++;
  } while (isLoaded);

  return result;}

  String get title {
    return txtTitle.text;
  }

  String get description {
    return txtDescription.text;
  }

  Key get imageKey {
    return new Key(uplProfileImg.dataset['id']);
  }

  Map<String, String> get endpoints {
    var result = new Map<String, String>();

    result.addEntries(getEndpoints().map((e) => e.toEntry()));

    return result;
  }

  Map<String, String> get codes {
    var result = new Map<String, String>();
    result.addEntries(getCodes().map((e) => e.toEntry()));

    return result;
  }

  Map<String, String> get terms {
    var result = new Map<String, String>();
    result.addEntries(getTerms().map((e) => e.toEntry()));

    return result;
  }
}