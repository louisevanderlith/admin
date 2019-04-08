import 'dart:html';
import 'formstate.dart';
import 'profileapi.dart';

class ProfileForm extends FormState {
  String _objKey;
  TextInputElement _name;
  TextAreaElement _description;
  EmailInputElement _email;
  TelephoneInputElement _phone;
  TextInputElement _url;
  FileUploadInputElement _image;
  
  ProfileForm(
      String idElem,
      String objKey,
      String nameElem,
      String descElem,
      String emailElem,
      String phoneElem,
      String urlElem,
      String imageElem,
      String submitBtn)
      : super(idElem, submitBtn) {
    _objKey = objKey;
    _name = querySelector(nameElem);
    _description = querySelector(descElem);
    _email = querySelector(emailElem);
    _phone = querySelector(phoneElem);
    _url = querySelector(urlElem);
    _image = querySelector(imageElem);

    querySelector(submitBtn).onClick.listen(onSend);
    registerValidation();
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

  void registerValidation() {
    _name.onBlur.listen((e) => {validate(e, _name)});
    _description.onBlur.listen((e) => {validateArea(e, _description)});
    _email.onBlur.listen((e) => {validate(e, _email)});
    _phone.onBlur.listen((e) => {validate(e, _phone)});
    _url.onBlur.listen((e) => {validate(e, _url)});
  }

  void validate(Event e, InputElement elem) {
    var elemValid = elem.checkValidity();

    if (!elemValid) {
      elem.setAttribute("invalid", "");
    } else {
      elem.removeAttribute("invalid");
    }

    elem.nextElementSibling.text = elem.validationMessage;

    super.disableSubmit(!super.isFormValid());
  }

  void validateArea(Event e, TextAreaElement elem) {
    var elemValid = elem.checkValidity();

    if (!elemValid) {
      elem.setAttribute("invalid", "");
    } else {
      elem.removeAttribute("invalid");
    }

    elem.nextElementSibling.text = elem.validationMessage;

    super.disableSubmit(!super.isFormValid());
  }

  void onSend(Event e) {
    if (isFormValid()) {
      disableSubmit(true);
      submitSend().then((obj) => {disableSubmit(false)});
    }
  }

  Future submitSend() async {
    var obj = {
      "Id": _objKey,
      "Title": name,
      "Description": description,
      "ContactEmail": email,
      "ContactPhone": phone,
      "URL": url,
      "ImageKey": imageKey,
    };

    return await updateProfile(obj);
  }
}
