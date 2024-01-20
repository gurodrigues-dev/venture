function logRequest(request)
    print("Método:", request.method)
    print("Rota:", request.route)
    print("Corpo:", request.body)
end