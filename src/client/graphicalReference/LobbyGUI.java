package graphicalReference;

import java.awt.*;
import java.io.File;
import java.io.IOException;
import java.util.ArrayList;

import javax.imageio.ImageIO;
import javax.swing.*;

import javax.swing.border.EmptyBorder;

import packageManaging.HostRoom;

public class LobbyGUI {
	public JPanel chatt = new JPanel();
	public JTextField field;
	public JTextPane chatPanel = new JTextPane();
	public JFrame lobby = new JFrame();
	public ChatGUI chatgui;
	public final String userName;
	public JButton joinButton;
	public JButton createButton;
	public JButton refreshButton;
	public JList<String> createList;
	public JList<HostRoom> joinList;
	public JTable jt;

	public JPanel westTopPanel;
	public JPanel eastTopPanel;

	public LobbyGUI(String usr) {
		this.userName = usr;

	}

	public void render(ArrayList<HostRoom> players, ArrayList<String> games) {

		lobby.setLayout(new BorderLayout());
		lobby.getContentPane().setBackground(Color.DARK_GRAY);

		lobby.setSize(1035, 720);
		lobby.setLocationRelativeTo(null);
		lobby.setResizable(false);
		lobby.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
		lobby.setBackground(Color.DARK_GRAY);

		JPanel topPanel = new JPanel();
		JPanel bottomPanel = new JPanel();

		eastTopPanel = new JPanel();
		westTopPanel = new JPanel();

		JPanel centerTopPanel = new JPanel();
		topPanel.setLayout(new BorderLayout());
		eastTopPanel.setLayout(new BorderLayout());
		westTopPanel.setLayout(new BorderLayout());
		refreshButton = new JButton();
		makeJButton(refreshButton, "resources/refreshButton.png",
				"refreshbutton");

		centerTopPanel.setLayout(new BorderLayout());
		refreshButton.setBackground(Color.DARK_GRAY);

		centerTopPanel.setBackground(Color.DARK_GRAY);
		topPanel.add(centerTopPanel, BorderLayout.CENTER);

		makePlayerList(eastTopPanel, players);
		makeGameList(westTopPanel, games);

		topPanel.add(westTopPanel, BorderLayout.WEST);
		topPanel.add(eastTopPanel, BorderLayout.EAST);
		topPanel.setBackground(Color.DARK_GRAY);

		lobby.add(topPanel, BorderLayout.NORTH);
		lobby.add(bottomPanel, BorderLayout.SOUTH);
		bottomPanel.add(chatt, "Center");
		chatgui = new ChatGUI(this.userName);
		chatgui.render(chatt, lobby);

		topPanel.setBorder(null);
		lobby.validate();
		lobby.setVisible(true);

	}

	public void makePlayerList(JPanel panel, ArrayList<HostRoom> L) {
		JPanel buttons = new JPanel(new GridLayout(1, 2));

		refreshButton.setBorder(null);

		joinButton = new JButton();
		panel.setBorder((new EmptyBorder(10, 30, 10, 30)));
		panel.setBackground(Color.DARK_GRAY);
		makeJButton(joinButton, "resources/JoinButton.png", "joinbutton");

		joinButton.setBackground(Color.DARK_GRAY);

		buttons.add(joinButton);

		buttons.add(refreshButton);
		buttons.setBackground(Color.DARK_GRAY);
		panel.add(buttons, BorderLayout.SOUTH);
		addArrayList(panel, L);

	}

	public void makeGameList(JPanel panel, ArrayList<String> L) {
		createButton = new JButton();

		makeJButton(createButton, "resources/CreateButton.png", "createbutton");
		panel.add(createButton, BorderLayout.SOUTH);
		addArrayListString(panel, L);

	}

	public void addArrayList(JPanel panel, ArrayList<HostRoom> L) {

		String[] title = { "Game", "Players", "Host", "RoomID" };

		jt = new JTable(new TableModel(makeNewTableArray(L), title));
		TableModel tm = (TableModel) jt.getModel();
		tm.fireTableStructureChanged();
		jt.setAutoCreateRowSorter(true);

		jt.setBackground(Color.DARK_GRAY.darker());
		jt.setForeground(Color.LIGHT_GRAY);
		jt.getColumnModel().getColumn(0).setPreferredWidth(150);
		jt.getColumnModel().getColumn(1).setPreferredWidth(25);
		jt.getColumnModel().getColumn(2).setPreferredWidth(100);
		jt.getColumnModel().getColumn(3).setPreferredWidth(25);
		JScrollPane jp = new JScrollPane(jt);

		jp.getViewport().setBackground(Color.DARK_GRAY.darker());
		jt.setRowSelectionAllowed(true);
		panel.add(jp, BorderLayout.CENTER);
		lobby.validate();

	}

	public String[][] makeNewTableArray(ArrayList<HostRoom> L) {
		String[][] array = new String[L.size()][4];
		for (int i = 0; i < array.length; i++) {
			array[i][0] = L.get(i).GameName;
			array[i][1] = Integer.toString(L.get(i).ClientCount) + "/"
					+ Integer.toString(L.get(i).MaxSize);
			array[i][2] = L.get(i).RoomName;
			array[i][3] = Integer.toString(L.get(i).RoomID);
		}

		return array;
	}

	public void updateJTable(ArrayList<HostRoom> L) {
		TableModel tm = (TableModel) jt.getModel();
		tm.removeAll();
		String[][] array = makeNewTableArray(L);
		tm.addAll(array);
		tm.fireTableDataChanged();
		jt.validate();
		eastTopPanel.validate();

	}

	public void addArrayListString(JPanel panel, ArrayList<String> L) {
		final JList<String> jList = getJListString(L);

		jList.setBackground(Color.DARK_GRAY.darker());
		jList.setForeground(Color.LIGHT_GRAY);
		panel.setBorder(new EmptyBorder(10, 50, 10, 50));
		panel.setBackground(Color.DARK_GRAY);

		JScrollPane js = new JScrollPane(jList);
		js.setHorizontalScrollBar(null);
		js.setBorder(null);
		panel.add(js, BorderLayout.CENTER);

		createList = jList;

	}

	public static JList<String> getJListString(ArrayList<String> list) {
		int size = list.size();
		final DefaultListModel<String> model = new DefaultListModel<String>();
		for (int i = 0; i < size; i++) {
			model.addElement(list.get(i));
		}
		return new JList<String>(model);
	}

	public void makeJButton(JButton jb, String src, String actionCommand) {
		File imageCheck = new File(src);
		try {
			Image img = ImageIO.read(imageCheck);
			jb.setIcon(new ImageIcon(img));
		} catch (IOException ex) {
			System.out.println("Hej");
		}
		jb.setMargin(new Insets(0, 0, 0, 0));
		jb.setBorder(null);
		jb.setActionCommand(actionCommand);

	}

}