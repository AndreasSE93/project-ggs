package clientHandlers;

import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.io.IOException;

import org.json.JSONException;
import org.json.JSONObject;

import packageManaging.CreateGameMessage;
import packageManaging.CreateGameMessageEncoder;
import packageManaging.JoinMessage;
import packageManaging.JoinMessageEncoder;
import packageManaging.LobbyMessageEncoder;
import packageManaging.LobbyServerMessage;
import packageManaging.Message;
import packageManaging.ChatMessageEncoder;
import graphicalReference.ChatGUI;
import graphicalReference.LobbyGUI;

import clientNetworking.NetManager;

public class LobbyHandler implements HandlerInterface,
		ActionListener {
	LobbyGUI lg;
	ChatGUI cg;
	LobbyMessageEncoder lme = new LobbyMessageEncoder();
	ChatMessageEncoder cme = new ChatMessageEncoder();
	JoinMessageEncoder jme = new JoinMessageEncoder();
	CreateGameMessageEncoder gme = new CreateGameMessageEncoder();
	NetManager network;
	boolean loop = true;
	int state;
	final String userName;
	public LobbyHandler(NetManager net, String username){
		this.network = net;
		this.userName = username;
	}
	
	public int initHandler() {

		String firstcall;
		LobbyServerMessage LM = null;
		try {
			firstcall = network.receiveMessage();
			LM = lme.decode(firstcall);
			
		} catch (IOException | JSONException e) {
			e.printStackTrace();
		}

		if (LM != null) {
			lg = new LobbyGUI(this.userName);
			lg.render(LM.UserList, LM.GameHost);
			lg.chatgui.field.addActionListener(this);
			lg.joinButton.addActionListener(this);
			lg.createButton.addActionListener(this);
			lg.refreshButton.addActionListener(this);
			
		} else {
			System.out.println("Connection failed");
		}

		while (loop) {
			try {
				String mess = recieveMessage();
			
				int id = retrieveId(mess);
				decodeAndRender(id, mess);
				
			} catch (IOException | JSONException e) {
				e.printStackTrace();
			}
		}
		
		//Change state!!!
		
		 lg.lobby.setVisible(false);
		 return state;
		 //TiarHandler th = new TiarHandler(network);
		 //th.init(lg.chatgui.userName);

	}

	public void actionPerformed(ActionEvent e) {
		String command = e.getActionCommand();
		String JSONtext = "";
		switch (command) {
		case "chatmessage":  //JTextField
			try {
				JSONtext = cme.encode(new Message(lg.chatgui.field.getText(), lg.chatgui.userName));
				lg.chatgui.field.setText("");
			} catch (JSONException e1) {
				e1.printStackTrace();
			}

			break;

		case "joinbutton":
			try {
				JSONtext = jme.encode(new JoinMessage((String)lg.joinList.getSelectedValue())); // Vad göra sen? skickat till serven att vi vill spela.
				} catch (JSONException e1) {
					// TODO Auto-generated catch block
					e1.printStackTrace();
				}
			state = 2;
			loop=false;


			break;
			
		case "createbutton":
			try {
				JSONtext = gme.encode(new CreateGameMessage((String) lg.createList.getSelectedValue()));
			} catch (JSONException e1) {
				// TODO Auto-generated catch block
				e1.printStackTrace();
			}
			
			break;
		case "refreshbutton":			//SKicka jag vill uppdatera spelarlistorna
			
			break;
		default:
			JSONtext = "";

		}
		sendMessage(JSONtext);

	}

	public void decodeAndRender(int id, String message) throws JSONException {
		switch (id) {
		case 100: // Message.java
			ChatMessageEncoder mh = new ChatMessageEncoder();
			Message chatMessage = mh.decode(message);
			lg.chatgui.chatUpdate(chatMessage.message, chatMessage.user);
			break;

		case 101: // Create session
			
			break;

		case 102: // Join session
			break;

		case 103: // LobbyClientMessage.java/LobbyServerMessage.java
			break;
		default:
			System.out.println("something went wrong!");
		}

	}

	public void sendMessage(String message) {
		network.send(message);
	}

	public String recieveMessage() throws IOException {
		return network.receiveMessage();

	}

	public int retrieveId(String mess) {
		try {
			JSONObject obj = new JSONObject(mess);
			return obj.getInt("PacketID");
		} catch (JSONException e) {
			
		}
		return 0;

	}
}
