import 'package:flutter/foundation.dart';
import 'package:auth0_flutter/auth0_flutter.dart';

class AuthModel extends ChangeNotifier {
  /// Internal, private state of the cart.
  Credentials? _credentials;

  /// The current total price of all items (assuming all items cost $42).
  String get userName => _credentials.user.name;
  String get userEmail => _credentials.user.email;
  String get userNickname => _credentials.user.nickname;

  void setCredentials(Credentials credentials) {
    _credentials = credentials;
    notifyListeners();
  }

  /// Adds [item] to cart. This and [removeAll] are the only ways to modify the
  // void add(Item item) {
  //   _items.add(item);
  //   // This call tells the widgets that are listening to this model to rebuild.
  //   notifyListeners();
  // }

  // /// Removes all items from the cart.
  // void removeAll() {
  //   _items.clear();
  //   // This call tells the widgets that are listening to this model to rebuild.
  //   notifyListeners();
  // }
}