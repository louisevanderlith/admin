import 'dart:html';

class VehicleExtra {
  List<String> get extraitems {
    final boxes = querySelectorAll("input[type=checkbox]");
    return boxes.map((e) => e.innerText);
  }
}
