import 'dart:convert';
import 'dart:html';
import 'formstate.dart';
import 'services/profileapi.dart';
import 'headerform.dart';
import 'models/headeritem.dart';
import 'portfolioform.dart';
import 'models/portfolioitem.dart';
import 'socialmediaform.dart';
import 'models/socialmediaitem.dart';
import 'services/uploadapi.dart';

class ProfileForm extends FormState {
  String _objKey;
  TextInputElement _name;
  TextAreaElement _description;
  EmailInputElement _email;
  TelephoneInputElement _phone;
  TextInputElement _url;
  FileUploadInputElement _image;

  HeaderForm _headers;
  PortfolioForm _portfolios;
  SocialmediaForm _socialmedia;

  ProfileForm(
      String idElem,
      String objKey,
      String nameElem,
      String descElem,
      String emailElem,
      String phoneElem,
      String urlElem,
      String imageElem,
      String frmHeader,
      String addHeader,
      String frmPortfolio,
      String addPortfolio,
      String frmSocialmedia,
      String addSocial,
      String submitBtn)
      : super(idElem, submitBtn) {
    _objKey = objKey;
    _name = querySelector(nameElem);
    _description = querySelector(descElem);
    _email = querySelector(emailElem);
    _phone = querySelector(phoneElem);
    _url = querySelector(urlElem);
    _image = querySelector(imageElem);

    _headers = new HeaderForm(frmHeader, submitBtn, addHeader);
    _socialmedia = new SocialmediaForm(frmSocialmedia, submitBtn, addSocial);
    _portfolios = new PortfolioForm(frmPortfolio, submitBtn, addPortfolio);

    querySelector(submitBtn).onClick.listen(onSend);

    _image.onChange.listen(uploadFile);
  }

  String get name {
    return _name.value;
  }

  String get description {
    return _description.value;
  }

  String get email {
    return _email.value;
  }

  String get phone {
    return _phone.value;
  }

  String get url {
    return _url.value;
  }

  String get imageKey {
    return _image.dataset["id"];
  }

  List<PortfolioItem> get portfolioItems {
    return _portfolios.items;
  }

  List<SocialmediaItem> get socialmediaItems {
    return _socialmedia.items;
  }

  List<HeaderItem> get headerItems {
    return _headers.items;
  }

  void onSend(MouseEvent e) async {
    if (isFormValid()) {
      disableSubmit(true);

      var resp = await updateProfile(_objKey, name, description, email, phone, url, imageKey, portfolioItems, socialmediaItems, headerItems);
      
    }
  }

  void onSuccess(String json) {
    print(jsonDecode(json));
  }
}
