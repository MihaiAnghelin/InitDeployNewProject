FROM mcr.microsoft.com/dotnet/sdk:{{DOTNET_VERSION}} AS build-env
WORKDIR /app

# Copy everything
COPY . ./
# Restore as distinct layers
RUN dotnet restore
# Build and publish a release
RUN dotnet publish -c Release -o out

# Build runtime image
FROM mcr.microsoft.com/dotnet/aspnet:{{DOTNET_VERSION}}
WORKDIR /app
COPY --from=build-env /app/out .


EXPOSE 5000
ENV ASPNETCORE_URLS=http://*:5000

ENTRYPOINT ["dotnet", "{{PROJECT_NAME}}.dll"]
